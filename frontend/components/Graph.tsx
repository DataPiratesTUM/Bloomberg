import { Line } from "@nivo/line";

export function Graph({ timeseries }: { timeseries: Timeseries }) {
  // Ziel ist https://nivo.rocks/storybook/?path=/docs/line--highlighting-negative-values
  let positive: Timeseries = [];
  let negative: Timeseries = [];

  const firstPrice = timeseries[0].price;
  const firstTimestamp = timeseries[0].timestamp;

  for (const point of timeseries) {
    point.timestamp -= firstTimestamp;
    if (point.price < firstPrice) {
      negative.push(point);
      // @ts-ignore
      positive.push({ timestamp: point.timestamp, price: null });
      break;
    }
    positive.push(point);
    // @ts-ignore
    negative.push({ timestamp: point.timestamp, price: null });
  }
  positive = positive.map((t) => {
    return { x: t.timestamp, y: t.price };
  });
  negative = negative.map((t) => {
    return { x: t.timestamp, y: t.price };
  });
  console.log(positive, negative);
  return (
    <Line
      width={700}
      height={400}
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
        type: "linear",
        min: -10000,
        max: 350000,
      }}
      yScale={{
        type: "linear",
        stacked: false,
        min: 0,
        max: 1000,
      }}
      enableArea={true}
      areaOpacity={0.07}
      enableSlices={false}
      useMesh={true}
      crosshairType="cross"
    />
  );
}
