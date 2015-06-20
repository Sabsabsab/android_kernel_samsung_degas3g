#!/bin/bash
rm -frv build
make clean && make mrproper
mkdir -pv build/out/modules build/out/dt_image
export ARCH=arm
export CROSS_COMPILE=~/android_prebuilts_toolchains/arm-cortex_a9-linux-gnueabi-linaro-4.9/bin/arm-cortex_a9-linux-gnueabihf-
export STRIP=~/android_prebuilts_toolchains/arm-cortex_a9-linux-gnueabi-linaro-4.9/bin/arm-cortex_a9-linux-gnueabihf-strip
make pxa1088_degas3g_eur_defconfig
# make menuconfig && wait
make CONFIG_NO_ERROR_ON_MISMATCH=y -j5 && make modules
./tools/dtbTool -o build/out/dt_image/boot.img-dt -p ./scripts/dtc/ ./arch/arm/boot/dts/
cp arch/arm/boot/*zImage build/out/boot.img-zImage
find -type f -name *.ko -exec cp {} build/out/modules/ \;
ls -al build/out/modules/
cd build/out/modules/
$STRIP --strip-unneeded *.ko
cd ../../../
ls -al build/out/modules/ build/out/dt_image/ build/out/
echo Done !
