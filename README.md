# The purpose of this fork

The original (logstash-forwarder)[https://github.com/elasticsearch/logstash-forwarder] seems to be broken on Windows platform:
 
 * it locks the log files, so they cannot be renamed by Log4j rolling appender
 * it does not detect file truncation very well
 * 
 Driskell did a very nice job fixing all the issues on Windows on his (fork)[https://github.com/driskell/logstash-forwarder], however his fixes were not merged into the Elastic Search logstash-forwarder (I don't know why). Driskell has abandoned his work on logstash-forwarder and started developing new (Log Courier)[https://github.com/driskell/log-courier].

I've made this fork because I wanted to use the Driskell's version of logstash-forwarder, since in fixes Windows issues and has a functional multiline codec (the multiline codec in Logstash is still problematic, because it waits for the next line to decide if the event has finished or not).
However, Driskell version of the logstash-forwared had a few problems which needed to be fixed:
 
 * compilation on Windows was broken
 * there was a bug in configuration file parsing when multiple file elements were provided

In my fork those issues are fixes.

There is still one more issue left which I did not track down yet - sometimes, when there are a lot of events coming and logstash-forwared looses connection to the logstash server it tries and fails to reconnect indefinitely. Restart of the service is required to recover. 
