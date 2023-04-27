# secret-files

Exploring ways to hide files inside of a go binary and only allow certain people to access them.

## PoC - secret-file.go

{add link here}

Using go:embed to load the /secrets folder on build.
The folder contains a file called `secret-key.txt` which is loaded into the binary.
To uncover the files the user has to pass a `-secret={key}` flag to the binary where the `{key}` matches the secret key hidden in the binary.

Proves the concept of hiding files in a binary but is not secure as the key is stored in the binary.
Someone could also just read the files out of the binary - as they are not encrypted.

### Problems from PoC

- The key can be read from the binary.
  - The key should not be in the binary, or required at all.
  - The key could be fetched at runtime from a secure location. If files are encrypted, we can eliminate the need for a password, since we'll need to pass a secret key to decrypt the files, anyway.
- Files themselves, are not encrypted, can be read from the binary.
  - The files should be encrypted, and we pass the decode secret when running the executable. If the key is correct, the files uncovered will be decrypted and readable.

## PoC / Solution #2 - secret-file-encrypted.go
