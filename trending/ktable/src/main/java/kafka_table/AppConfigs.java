package kafka_table;

public class AppConfigs {
    final static String applicationID = "StreamingTable";
    final static String bootstrapServers = "localhost:9092";
    final static String topicName = "top1";
    final static String stateStoreLocation = "tmp/state-store";
    final static String stateStoreName = "kt01-store";
    final static String queryServerHost = "localhost";
    final static int queryServerPort = 7010;
    final static String to_topic = "top2";
}
