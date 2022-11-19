import { Line } from "@nivo/line";

export function Graph({ timeseries }: { timeseries: Timeseries }) {
  let positive = [];
  let negative = [];

  const firstPrice = timeseries[0].price;

  for (const point of timeseries) {
    if (point.price < firstPrice) {
      negative.push(point);
      break;
    }
    positive.push(point);
  }
  positive = positive.map((t) => {
    const date = new Date(t.timestamp);
    console.log(date.toISOString());
    return { x: date.toISOString(), y: t.price };
  });
  negative = negative.map((t) => {
    const date = new Date(t.timestamp);
    console.log(date.toISOString());
    return { x: date.toISOString(), y: t.price };
  });
  console.log(positive, negative);
  return (
    <Line
      width={500}
      height={300}
      data={[
        {
          id: "positive :)",
          data: positive,
        },
        {
          id: "negative :(",
          data: negative,
        },
      ]}
      curve="monotoneX"
      enablePointLabel={true}
      pointSize={14}
      pointBorderWidth={1}
      pointBorderColor={{
        from: "color",
        modifiers: [["darker", 0.3]],
      }}
      pointLabelYOffset={-20}
      enableGridX={false}
      colors={["rgb(97, 205, 187)", "rgb(244, 117, 96)"]}
      xScale={{
        type: "time",
        format: "%Y-%m-%dT%H:%M:%S.%L%Z",
      }}
      xFormat="time:%Y-%m-%dT%H:%M"
      yScale={{
        type: "linear",
        stacked: false,
      }}
      enableArea={true}
      areaOpacity={0.07}
      enableSlices={false}
      useMesh={true}
      crosshairType="cross"
    />
  );
}
