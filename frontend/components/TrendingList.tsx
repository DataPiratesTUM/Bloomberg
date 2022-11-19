import { TrendingElement } from "./TrendingElement";

export function TrendingList({ trendingList }: { trendingList: TrendingList }) {
  return (
    <div className="border-4 border-gray-800 p-0.5 rounded  place-self-center">
      <p>Trending</p>
      {trendingList.trendings.map((trending) => (
        <TrendingElement
          key={trending.security_id}
          trendingElement={trending}
        />
      ))}
    </div>
  );
}
