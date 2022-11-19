import Link from "next/link";

export function TrendingElement({
  trendingElement,
}: {
  trendingElement: TrendingSec;
}) {
  return (
    <Link
      href={"/securities/" + trendingElement.security_id}
      key={trendingElement.security_id}
    >
      <div className=" m-2 border shadow rounded my-2 p-4 flex justify-between">
        <p>{trendingElement.title}</p>
      </div>
    </Link>
  );
}
