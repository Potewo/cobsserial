#include "Arduino.h"
#include <PacketSerial.h>

PacketSerial myPacketSerial;

void onPacketReceived(const uint8_t* buffer, size_t size) {
  /* uint8_t tempBuffer[size]; */
  /* memcpy(tempBuffer, buffer, size); */
  delay(100);
  myPacketSerial.send(buffer, size);
}

void setup() {
  myPacketSerial.begin(9600);
  myPacketSerial.setPacketHandler(&onPacketReceived);
}

void loop() {
  myPacketSerial.update();
  if (myPacketSerial.overflow()) {
  }
}

