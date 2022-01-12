TEMPLATE = app
TARGET = refresh
INCLUDEPATH += .

# Input
SOURCES += cli.cpp refresh.cpp
LIBS += -L/tc/arm-nickel-linux-gnueabihf/arm-nickel-linux-gnueabihf/sysroot/usr/lib -ldl