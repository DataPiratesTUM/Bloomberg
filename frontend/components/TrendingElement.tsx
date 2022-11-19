export function TrendingElement({
  trendingElement,
}: {
  trendingElement: TrendingSec;
}) {
  return (
    <div className="border-2 m-2 p-0.5">
      <p>{trendingElement.title}</p>
    </div>
  );
}
