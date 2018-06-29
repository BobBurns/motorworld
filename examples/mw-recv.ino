// template to add fuctionality to web app controlled 
// particle device

int led = D7;  // The on-board LED


void setup() {
  pinMode(led, OUTPUT);
  Particle.subscribe("indata", myHandler);
}

void myHandler(const char *event, const char *data)
{
         
  if (String(data).compareTo("on") == 0) {
    digitalWrite(led, HIGH);
  } else if (String(data).compareTo("off") == 0) {
    digitalWrite(led, LOW);
  } else if (String(data).compareTo("rev") == 0) {
      // reverse motor code
  } else if (String(data).compareTo("fast") == 0) {
      // fast motor code
  } else if (String(data).compareTo("med") == 0) {
      // med motor code
  } else if (String(data).compareTo("slow") == 0) {
      // slow motor code
  }
}

void loop() {
  String temp = String(random(60, 80));
  Particle.publish("outdata", temp, PRIVATE);
  delay(60000);               // Wait for 30 seconds

}
