# Membuffers

## Wire Format

* Alignment is to size of elements but no more than 4

### Vectors

#### message

```c
uint32 size                     // total size in bytes of the entire message [aligned to 4]

// actual fields                // [aligned to 4]
```

#### bytes / string

```c
uint32 size                     // total size in bytes of the array [aligned to 4]

// actual bytes
```

#### repeated uint8 / uint16 / uint32

```c
uint32 size                     // total size in bytes of the array [aligned to 4]

// actual bytes                 // aligned to size of element
```

#### repeated uint64

```c
uint32 size                     // total size in bytes of the array [aligned to 4]

// actual bytes                 // [aligned to 4]
```

#### repeated message

```c
uint32 size                     // total size in bytes of the array [aligned to 4]

// actual messages              // [aligned to 4
```