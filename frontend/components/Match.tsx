export function Match({ match }: { match: Match }) {
  return (
    <div key={match.created}>
      <div className=" m-2 border shadow rounded my-2 p-4 flex justify-between">
        <p>
          {match.security}: {match.quantity} @ {match.price}
        </p>
      </div>
    </div>
  );
}
