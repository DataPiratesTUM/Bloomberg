import { TrendingElement } from "./TrendingElement";

export function TrendingList({ trendingList }: { trendingList: TrendingList }) {
  return (
    <div className=" p-0.5 rounded  place-self-center">
      <h3 className="text-4xl font-bold tracking-tight  sm:text-5xl">
        Trending
      </h3>
      {trendingList.trendings.map((trending) => (
        <TrendingElement
          key={trending.security_id}
          trendingElement={trending}
        />
      ))}
    </div>
  );
}
