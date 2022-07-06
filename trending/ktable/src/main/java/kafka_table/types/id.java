package kafka_table.types;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import org.apache.commons.lang3.builder.ToStringBuilder;

import java.text.SimpleDateFormat;
import java.time.LocalDate;
import java.util.Date;

@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder({
        "Id",
        "Views"
})
public class id {
    @JsonProperty("Id")
    private String Id;
    @JsonProperty("Views")
    private Integer Views;

    @JsonProperty("UpDate")
    private String UpDate;

    @JsonProperty("Days")
    private long Days;

    @JsonProperty("Id")
    public  String getId(){
        return Id;
    }

    @JsonProperty("Views")
    public Integer getViews(){
        return Views;
    }

    @JsonProperty("UpDate")
    public String getUpDate(){
        return UpDate;
    }

    @JsonProperty("Days")
    public  long getDays(){
        return Days;
    }

    @JsonProperty("Id")
    public void setId(String Id) {
        this.Id = Id;
    }
    @JsonProperty("Views")
    public void setViews( Integer Views){
        this.Views = Views;
    }

    @JsonProperty("UpDate")
    public void setUpDate(String UpDate){
        this.UpDate = UpDate;
    }

    @JsonProperty("Days")
    public void setDays(long Days) {
        this.Days = Days;
    }

    public id withId(String Id){
        this.Id = Id;
        return this;
    }

    public id withViews( Integer Views) {
        this.Views = Views;
        return this;
    }
    public id withUpDate( String UpDate) {
        this.UpDate = UpDate;
        return this;
    }

    public id withDays( long Days) {
        this.Days = Days;
        return this;
    }
    @Override
    public String toString() {
        return new ToStringBuilder(this).
                append("Id",Id).
                append("Views", Views).
                append("DelDate", UpDate).
                append("Days", Days).
                toString();
    }
}
