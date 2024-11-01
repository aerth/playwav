# playwav

## play a .wav file

  * stuck with FormatS16LE for now (Little Endian 16, 2 bytes per sample)
  * compilation adds requirements: CGO and some c libraries (ALSA)

## playwav command
  ```go get -v github.com/aerth/playwav/cmd/playwav```
  
if not build because cant found 'alsa/a ... .h" 
try it:
```
sudo apt-get install libasound2-dev
```

## playwav library

```
import "github.com/aerth/playwav"

// from file
playwav.FromFile("example.wav")

```

## cool test with playwav command

```
for i in $(find /usr/ -name '*.wav'); do ./playwav $i; done
```

## output of ldd for playwav command:

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
