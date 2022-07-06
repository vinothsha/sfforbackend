package kafka_table.AppSerdes;

import kafka_table.types.id;
import org.apache.kafka.common.serialization.Serde;
import org.apache.kafka.common.serialization.Serdes;

import java.util.HashMap;
import java.util.Map;

public class AppSerdes extends Serdes {

    static public final class idSerde extends WrapperSerde<id> {
        public idSerde() {
            super(new JsonSerializer<>(), new JsonDeserializer<>(id.class));
        }
    }

    static public Serde<id> id() {
        idSerde serde = new idSerde();

        Map<String, Object> serdeConfigs = new HashMap<>();
        serdeConfigs.put(JsonDeserializer.VALUE_CLASS_NAME_CONFIG,id.class);
        serde.configure(serdeConfigs, false);

        return serde;
    }
}
