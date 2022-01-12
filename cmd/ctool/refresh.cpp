#include <QCoreApplication>
#include <QTimer>

#include <cstdlib>

#include "cli.h"

int main(int argc, char *argv[])
{
    QCoreApplication app(argc, argv);

    int timeOutSec = 0;
    if (argc == 2) {
        timeOutSec = atoi(argv[1]);
    }

    Cli cli(&app, timeOutSec);
    QTimer::singleShot(0, &cli, SLOT(start()));

    return app.exec();
}

