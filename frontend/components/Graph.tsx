import { Line, ResponsiveLine } from "@nivo/line";
import { format } from "date-fns";

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
  if (timeseries.length === 0) {
    return (
      <div className="mt-64 text-lg text-center font-mono">
        Once there are successful order matches, a detailed price graph will be
        shown
      </div>
    );
  }

  timeseries.reduce<Timeseries>((acc, cur) => {
    if (
      acc[acc.length - 1] &&
      acc[acc.length - 1].timestamp !== cur.timestamp
    ) {
      acc.push(cur);
    }
    return acc;
  }, []);
  // Ziel ist https://nivo.rocks/storybook/?path=/docs/line--highlighting-negative-values
  let positive: Dot[] = [];
  let negative: Dot[] = [];

  const firstPrice = timeseries[0].price;
  const firstTimestamp = timeseries[0].timestamp;

  timeseries.forEach((point, i, series) => {
    const date = format(new Date(point.timestamp), "yyyy-MM-dd:HH:mm:ss");
    const dot = {
      x: date,
      y: point.price,
    };
    const nullDot = { ...dot, y: null };
    if (point.price < firstPrice) {
      if (series[i - 1].price > firstPrice) {
        // Wendepunkt
        const previousPoint = series[i - 1];
        const previousDate = format(
          new Date(previousPoint.timestamp),
          "yyyy-MM-dd:HH:mm:ss"
        );
        negative.push({
          x: previousDate,
          y: previousPoint.price,
        });
      }
      // Order of insertion matters
      negative.push(dot);
      positive.push(nullDot);

      if (series[i + 1] && series[i + 1].price >= firstPrice) {
        const nextPoint = series[i + 1];
        const nextDate = format(
          new Date(nextPoint.timestamp),
          "yyyy-MM-dd:HH:mm:ss"
        );
        negative.push({
          x: nextDate,
          y: nextPoint.price,
        });
      }
    } else {
      negative.push(nullDot);
      positive.push(dot);
    }
  });

  return (
    <ResponsiveLine
      margin={{ top: 40, right: 40, bottom: 60, left: 80 }}
      animate={true}
      height={500}
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
        format: "%Y-%m-%d:%H:%M:%S",
        nice: true,
      }}
      xFormat="time:%Y-%m-%d:H"
      yScale={{
        type: "linear",
        stacked: false,
        min: 0,
        max: "auto",
        nice: true,
      }}
      axisLeft={{
        legend: "Portfolio value in €",
        legendOffset: 12,
      }}
      axisBottom={{
        format: "%b %d",
        //tickValues: "every 2 days",
        // tickRotation: -90,
        legend: "time scale",
        legendOffset: -12,
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
