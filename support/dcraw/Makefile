
all: native windows

native:
	gcc -o dcraw -O4 dcraw.c -lm -DNODEPS

# mingw-w64 / mingw32
windows:
	i586-mingw32msvc-gcc   -o dcraw.win32.exe -O4 dcraw.c -lm -DNODEPS
	x86_64-w64-mingw32-gcc -o dcraw.win64.exe -O4 dcraw.c -lm -DNODEPS

