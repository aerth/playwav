# playwav

compilation requires CGO and some c libraries (ALSA)

## output of ldd:

```
    linux-vdso.so.1 (0x00007fff27fcf000)
    libasound.so.2 => /usr/lib64/libasound.so.2 (0x00007f418a496000)
    libpthread.so.0 => /lib64/libpthread.so.0 (0x00007f418a278000)
    libc.so.6 => /lib64/libc.so.6 (0x00007f4189eb2000)
    libm.so.6 => /lib64/libm.so.6 (0x00007f4189ba9000)
    libdl.so.2 => /lib64/libdl.so.2 (0x00007f41899a5000)
    librt.so.1 => /lib64/librt.so.1 (0x00007f418979b000)
    /lib64/ld-linux-x86-64.so.2 (0x000055ab41acf000)
```
