package imwsutil

import (
	"encoding/binary"
	"github.com/gobwas/pool/pbytes"
	"io"
)

// CipherReader implements io.Reader that applies xor-cipher to the bytes read
// from source.
// It could help to unmask WebSocket frame payload on the fly.
type CipherReader struct {
	r    io.Reader
	mask [4]byte
	pos  int
}

// NewCipherReader creates xor-cipher reader from r with given mask.
func NewCipherReader(r io.Reader, mask [4]byte) *CipherReader {
	return &CipherReader{r, mask, 0}
}

// Reset resets CipherReader to read from r with given mask.
func (c *CipherReader) Reset(r io.Reader, mask [4]byte) {
	c.r = r
	c.mask = mask
	c.pos = 0
}

// Read implements io.Reader interface. It applies mask given during
// initialization to every read byte.
func (c *CipherReader) Read(p []byte) (n int, err error) {
	n, err = c.r.Read(p)
	Cipher(p[:n], c.mask, c.pos)
	c.pos += n
	return
}

// CipherWriter implements io.Writer that applies xor-cipher to the bytes
// written to the destination writer. It does not modify the original bytes.
type CipherWriter struct {
	w    io.Writer
	mask [4]byte
	pos  int
}

// NewCipherWriter creates xor-cipher writer to w with given mask.
func NewCipherWriter(w io.Writer, mask [4]byte) *CipherWriter {
	return &CipherWriter{w, mask, 0}
}

// Reset reset CipherWriter to write to w with given mask.
func (c *CipherWriter) Reset(w io.Writer, mask [4]byte) {
	c.w = w
	c.mask = mask
	c.pos = 0
}

// Write implements io.Writer interface. It applies masking during
// initialization to every sent byte. It does not modify original slice.
func (c *CipherWriter) Write(p []byte) (n int, err error) {
	cp := pbytes.GetLen(len(p))
	defer pbytes.Put(cp)

	copy(cp, p)
	Cipher(cp, c.mask, c.pos)
	n, err = c.w.Write(cp)
	c.pos += n

	return
}

func Cipher(payload []byte, mask [4]byte, offset int) {
	n := len(payload)
	if n < 8 {
		for i := 0; i < n; i++ {
			payload[i] ^= mask[(offset+i)%4]
		}
		return
	}

	// Calculate position in mask due to previously processed bytes number.
	mpos := offset % 4
	// Count number of bytes will processed one by one from the beginning of payload.
	ln := remain[mpos]
	// Count number of bytes will processed one by one from the end of payload.
	// This is done to process payload by 8 bytes in each iteration of main loop.
	rn := (n - ln) % 8

	for i := 0; i < ln; i++ {
		payload[i] ^= mask[(mpos+i)%4]
	}
	for i := n - rn; i < n; i++ {
		payload[i] ^= mask[(mpos+i)%4]
	}

	// NOTE: we use here binary.LittleEndian regardless of what is real
	// endianness on machine is. To do so, we have to use binary.LittleEndian in
	// the masking loop below as well.
	var (
		m  = binary.LittleEndian.Uint32(mask[:])
		m2 = uint64(m)<<32 | uint64(m)
	)
	// Skip already processed right part.
	// Get number of uint64 parts remaining to process.
	n = (n - ln - rn) >> 3
	for i := 0; i < n; i++ {
		var (
			j     = ln + (i << 3)
			chunk = payload[j : j+8]
		)
		p := binary.LittleEndian.Uint64(chunk)
		p = p ^ m2
		binary.LittleEndian.PutUint64(chunk, p)
	}
}

// remain maps position in masking key [0,4) to number
// of bytes that need to be processed manually inside Cipher().
var remain = [4]int{0, 3, 2, 1}
