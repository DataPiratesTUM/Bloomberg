import { Line } from "@nivo/line";

interface Dot {
  x: number;
  y: number | null;
}

export function Graph({ timeseries }: { timeseries: Timeseries }) {
  // Ziel ist https://nivo.rocks/storybook/?path=/docs/line--highlighting-negative-values
  let positive: Dot[] = [];
  let negative: Dot[] = [];

  const firstPrice = timeseries[0].price;
  const firstTimestamp = timeseries[0].timestamp;

  timeseries.forEach((point, i, series) => {
    let dot = { x: point.timestamp - firstTimestamp, y: point.price };
    let nullDot = { ...dot, y: null };
    if (point.price < firstPrice) {
      if (series[i - 1].price > firstPrice) {
        // Wendepunkt
        const previousPoint = series[i - 1];
        negative.push({ x: previousPoint.timestamp - firstTimestamp, y: previousPoint.price });
      }
      // Order of insertion matters
      negative.push(dot);
      positive.push(nullDot);

      if (series[i + 1] && series[i + 1].price >= firstPrice) {
        const nextPoint = series[i + 1];
        negative.push({ x: nextPoint.timestamp - firstTimestamp, y: nextPoint.price });
      }
    } else {
      negative.push(nullDot);
      positive.push(dot);
    }
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
      curve="natural"
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
