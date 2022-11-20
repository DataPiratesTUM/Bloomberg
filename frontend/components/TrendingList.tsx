import Link from "next/link";
import { useQuery } from "@tanstack/react-query";

export function TrendingList() {
  const query = useQuery(["securities"], async () => {
    const res = await fetch("https://organisation.ban.app/security/search/title?query=");
    const securities: SecurityOverview[] = await res.json();
    return securities;
  });

  return (
    <div className=" p-0.5 rounded  place-self-center">
      <h3 className="text-4xl font-bold tracking-tight mb-8  sm:text-5xl">Trending</h3>
      {query.isLoading
        ? "Loading..."
        : query.isError
        ? "Error!"
        : query.data
        ? query.data.map((s) => {
            return (
              <Link href={"/securities/" + s.Id} key={s.Id}>
                <div className=" m-2 border shadow rounded my-2 p-4 flex justify-between">
                  <p>{s.Name}</p>
                </div>
              </Link>
            );
          })
        : null}
    </div>
  );
}
