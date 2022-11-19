declare type Timeseries = {
  timestamp: number;
  price: number;
}[];

declare interface Order {
  security: string;
  quantity: number;
  price: number;
  side: "buy" | "sell";
}

declare interface Match {
  created: number;
  price: number;
  quantity: number;
  security: string;
}

declare interface Security {
  security_id: string;
  price: number;
  title: string;
  creator: {
    name: string;
    organisation_id: string;
  };
  description: string;
  funding_amount: number;
  creation_date: number;
  funding_date: number | null;
  ttl_phase_one: number;
  ttl_phase_two: number;
  timeseries: Timeseries;
  orders: Order[];
}

declare interface User {
  user_id: string;
  name: string;
  balance: number;
  securities: {
    id: string;
    name: string;
    qty: number;
    price: number;
    price_bought: number;
  }[];
  timeseries: Timeseries;
}

declare interface TrendingList {
  trendings: TrendingSec[];
}

declare interface TrendingSec {
  security_id: string;
  title: string;
}
