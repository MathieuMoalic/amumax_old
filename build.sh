#!/bin/bash

# The cuda versions against which we will compile mumax3
CUDAVERSION=10.1
export NVCC_CCBIN=/usr/bin/gcc-8
export CUDA_CC="   30 32 35 37 50 52 53 60 61 62 70 72 75"

# The final location of the mumax3 executables and libs
MUMAX3UNAME=amumax
BUILDDIR=./build/${MUMAX3UNAME}
rm -rf $BUILDDIR
mkdir -p $BUILDDIR

# The path for shared libraries (relative to the build directory)
RPATH=lib 
mkdir -p $BUILDDIR/$RPATH

# We overwrite the CGO Flags to make sure that it is compiled against $CUDAVERSION
export LD_LIBRARY_PATH=/usr/lib/x86_64-linux-gnu/:$LD_LIBRARY_PATH
export CGO_LDFLAGS="-lcufft -lcurand -lcuda -L /usr/lib/x86_64-linux-gnu/ -Wl,-rpath -Wl,\$ORIGIN/$RPATH"
export CGO_CFLAGS="-I /usr/include/include"

# (Re)build everything
# (make realclean && make -j 11 || exit 1)
make -j 11
# Copy the executable and the cuda libraries to the output directory
cp $GOPATH/bin/amumax $BUILDDIR 
cp $( ldd ${BUILDDIR}/amumax | grep libcufft | awk '{print $3}' ) ${BUILDDIR}/${RPATH}
cp $( ldd ${BUILDDIR}/amumax | grep libcurand | awk '{print $3}' ) ${BUILDDIR}/${RPATH}

(cd build && tar -czf ${MUMAX3UNAME}.tar.gz ${MUMAX3UNAME})
