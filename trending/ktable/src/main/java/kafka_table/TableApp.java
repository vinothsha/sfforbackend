package kafka_table;
import kafka_table.AppSerdes.AppSerdes;
import kafka_table.types.id;
import org.apache.kafka.common.utils.Bytes;
import org.apache.kafka.streams.kstream.*;
import org.apache.kafka.streams.state.KeyValueStore;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import java.util.Properties;
import org.apache.kafka.streams.KafkaStreams;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.StreamsConfig;
import org.apache.kafka.common.serialization.Serdes;
import java.util.*;
import java.text.*;
import java.time.temporal.ChronoUnit;
import java.time.LocalDate;

public class TableApp extends java.lang.Object{
    private final static Logger logger = LogManager.getLogger();
   public static void main(String[] args) {
       Date dNow = new Date();
       SimpleDateFormat ft =
               new SimpleDateFormat ("yyyy-MM-dd");
       LocalDate CurrDate = LocalDate.parse(ft.format(dNow));
       final Properties props = new Properties();
        props.put(StreamsConfig.APPLICATION_ID_CONFIG,AppConfigs.applicationID);
        props.put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG,AppConfigs.bootstrapServers);
        props.put(StreamsConfig.STATE_DIR_CONFIG,AppConfigs.stateStoreLocation);
//        props.put(StreamsConfig.DEFAULT_KEY_SERDE_CLASS_CONFIG,Serdes.String().getClass().getName());
//        props.put(StreamsConfig.DEFAULT_VALUE_SERDE_CLASS_CONFIG,Serdes.String().getClass().getName());

        final StreamsBuilder builder = new StreamsBuilder();
        KStream<String, id> KS0 = builder.<String, id>stream(AppConfigs.topicName,
                Consumed.with(Serdes.String(), AppSerdes.id()));
        KStream<String, id> KS1 = KS0.selectKey((k, v) -> v.getId());
        KS1.foreach((k,v)-> v.setDays(ChronoUnit.DAYS.between(LocalDate.parse(v.getUpDate()),CurrDate)));
        KS1.foreach((k,v) -> v.setViews(v.getViews()+1));
        KTable<String, id> KT0 = KS1.<String,String>toTable(Materialized.<String,id, KeyValueStore<Bytes, byte[]>>as("new_table")
                .withKeySerde(Serdes.String())
                .withValueSerde(AppSerdes.id()));
        KTable<String,id> KT1 = KT0.filter((k,v) -> v.getDays() <= 15);
        KT1.toStream().to(AppConfigs.to_topic, Produced.with(Serdes.String(),AppSerdes.id()));
        KT1.toStream().print(Printed.toSysOut());
        KafkaStreams streams = new KafkaStreams(builder.build(),props);
//        QueryServer queryServer = new QueryServer(streams,AppConfigs.queryServerHost,AppConfigs.queryServerPort);
//        streams.setStateListener((newstate,oldstate) -> {
//            logger.info("State changing form " + oldstate +" to "+ newstate);
//            queryServer.setActive(newstate == KafkaStreams.State.RUNNING && oldstate  == KafkaStreams.State.REBALANCING);
//        });

        streams.start();
//        queryServer.start();

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            logger.info("Shutting down server");
//            queryServer.stop();
            streams.close();
        }));
    }
}
