#ifndef CLI_H
#define CLI_H

#include <QObject>

class Cli : public QObject {
    public:
        Cli(QObject* parent, int sec);
        int immportBooks();

    public Q_SLOTS:
        void start();
        void handleTimeout();
    private:
        int timeoutSec;
};

#endif /*CLI_H*/