CC = go build
TARGET = imdb
PREFIX ?= /usr/local

$(TARGET): $(TARGET).go
	$(CC) -o $(TARGET) $(TARGET).go

clean:
	rm -f $(TARGET)

uninstall:
	rm -f $(PREFIX)/bin/$(TARGET)

install: $(TARGET)
	install -m 0755 $(TARGET) $(PREFIX)/bin/$(TARGET)

all: $(TARGET)