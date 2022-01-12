#ifndef CLI_H
#define CLI_H

#include <QObject>

class Cli : public QObject {
    public:
        Cli(QObject* parent);
        int immportBooks();

    public Q_SLOTS:
        void start();
};

#endif /*CLI_H*/