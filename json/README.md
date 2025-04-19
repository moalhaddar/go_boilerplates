# JSON

There are two ways of handling JSON:
- json encoder decoder
- json marshal unmarshal

The encoder/decoder deals with io.Writer/Reader, hence it operates on streams without waiting for the full string to be in memory.

The marshal/unmarshal deals with byte arrays, expecting the whole payload to be in memory first.

The json struct tags are documented under json.Marshal