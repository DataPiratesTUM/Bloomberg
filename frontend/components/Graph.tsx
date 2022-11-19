import { Line, ResponsiveLine } from "@nivo/line";

interface Dot {
  x: string;
  y: number | null;
}
// @ts-ignore
const CustomSymbol = ({ size, color, borderWidth, borderColor }) => (
  <g>
    <circle
      fill="#fff"
      r={size / 2}
      strokeWidth={borderWidth}
      stroke={borderColor}
    />
    <circle
      r={size / 5}
      strokeWidth={borderWidth}
      stroke={borderColor}
      fill={color}
      fillOpacity={0.35}
    />
  </g>
);

export function Graph({ timeseries }: { timeseries: Timeseries }) {
  return <></>;
  // Ziel ist https://nivo.rocks/storybook/?path=/docs/line--highlighting-negative-values
  let positive: Dot[] = [];
  let negative: Dot[] = [];

  const firstPrice = timeseries[0].price;
  const firstTimestamp = timeseries[0].timestamp;

  timeseries.forEach((point, i, series) => {
    let dot = { x: new Date(point.timestamp).toISOString(), y: point.price };
    // let nullDot = { ...dot, y: null };
    if (point.price < firstPrice) {
      if (series[i - 1].price > firstPrice) {
        // Wendepunkt
        const previousPoint = series[i - 1];
        negative.push({
          x: new Date(previousPoint.timestamp).toISOString(),
          y: previousPoint.price,
        });
      }
      // Order of insertion matters
      negative.push(dot);
      //  positive.push(nullDot);

      if (series[i + 1] && series[i + 1].price >= firstPrice) {
        const nextPoint = series[i + 1];
        negative.push({
          x: new Date(nextPoint.timestamp).toISOString(),
          y: nextPoint.price,
        });
      }
    } else {
      //  negative.push(nullDot);
      positive.push(dot);
    }
  });

  console.log(positive, negative);
  return (
    <ResponsiveLine
      margin={{ top: 20, right: 20, bottom: 60, left: 80 }}
      animate={true}
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
      curve="linear"
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
        format: "%Y-%m-%dT%H:%M:%SZ",
      }}
      xFormat="time:%Y-%m-%dT%H:%M:%SZ"
      yScale={{
        type: "linear",
        stacked: false,
        min: 0,
        max: "auto",
      }}
      axisLeft={{
        legend: "Portfolio value in €",
        legendOffset: 12,
      }}
      axisBottom={{
        display: false,
        legend: "Time",
        legendOffset: -12,
        format: "%b %d",
      }}
      pointSymbol={CustomSymbol}
      enableArea={true}
      areaOpacity={0.07}
      enableSlices={false}
      useMesh={true}
      crosshairType="cross"
    />
  );
}
