# Membuffers

## Wire Format

### Vectors

#### message

```c
uint32 size                     // total size in bytes of the entire message [aligned to 4]

// actual fields                // [aligned to 8]
```

#### bytes / string

```c
uint32 size                     // total size in bytes of the array [aligned to 4]

// actual bytes
```

#### repeated uint16 / uint32 / uint64

```c
uint32 size                     // total size in bytes of the array [aligned to 4]

// actual bytes                 // aligned to size of element
```

#### repeated message

```c
uint32 size                     // total size in bytes of the array [aligned to 4]

// actual messages              // [aligned to 4
```