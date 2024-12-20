package day09

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
)

type Chunk struct {
	id           int
	fileSizeMax  int
	fileSizeLeft int
	freeSpace    int
	freeContents []int
}

func Solve(path string, partOne bool) int {
	chunks := parse(path)
	if partOne {
		return allocateSingle(chunks)
	}
	return allocateGroup(chunks)
}

func allocateGroup(chunks []Chunk) int {
	for fromIdx := len(chunks) - 1; fromIdx >= 0; fromIdx-- {
		from := &chunks[fromIdx]

		for intoIndex := 0; intoIndex < fromIdx; intoIndex++ {
			into := &chunks[intoIndex]

			if from == into {
				continue
			}

			if from.fileSizeLeft == 0 {
				continue
			}

			if into.freeSpace >= from.fileSizeMax {
				assert.Equal(from.fileSizeMax, from.fileSizeLeft, "source file size must be equal to used size for p2")
				into.freeSpace -= from.fileSizeMax
				for range from.fileSizeMax {
					into.freeContents = append(into.freeContents, from.id)
				}
				from.fileSizeLeft = 0
			}
		}

	}

	return checksum(chunks)
}

func allocateSingle(chunks []Chunk) int {
	take := func(chunk *Chunk) (int, bool) {
		if chunk.fileSizeLeft == 0 {
			return 0, false
		}
		chunk.fileSizeLeft -= 1
		return chunk.id, true

	}

	allocate := func(chunk *Chunk, value int) {
		assert.True(chunk.freeSpace > 0, "can not allocate, no free space remains", "chunk", chunk)
		chunk.freeSpace -= 1
		chunk.freeContents = append(chunk.freeContents, value)
	}

	for idx := 0; idx < len(chunks); idx++ {
		chunk := &chunks[idx]
		for range chunk.freeSpace {
			next := -1
			for i := len(chunks) - 1; i > idx; i-- {
				if candidate, ok := take(&chunks[i]); ok {
					assert.True(chunks[i].freeSpace >= 0, "cannot have negative free space after allocating")
					next = candidate
					break
				}
			}
			if next == -1 {
				break
			}

			assert.True(chunk.freeSpace >= 0, "must have positive free space")
			allocate(chunk, next)
		}
	}
	return checksum(chunks)
}

func checksum(chunks []Chunk) int {
	checksum := 0
	offset := 0

	for _, chunk := range chunks {
		for idx := range chunk.fileSizeLeft {
			checksum += (idx + offset) * chunk.id
		}
		for idx, v := range chunk.freeContents {
			checksum += (idx + offset + chunk.fileSizeMax) * v
		}
		offset += chunk.fileSizeMax + len(chunk.freeContents) + chunk.freeSpace
	}

	return checksum

}

func parse(path string) []Chunk {
	lines := utils.MustReadInput(path)
	assert.Equal(1, len(lines), "expected a single line of input")
	line := lines[0]
	chunks := make([]Chunk, 0)

	for i := 0; i < len(line); i = i + 2 {
		free := utils.MustAtoi(string(line[i+1]))
		size := utils.MustAtoi(string(line[i]))
		chunks = append(chunks, Chunk{
			id:           i / 2,
			fileSizeLeft: size,
			fileSizeMax:  size,
			freeSpace:    free,
			freeContents: make([]int, 0, free),
		})
	}
	return chunks
}
