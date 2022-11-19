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
  creator: string;
  description: string;
  fundingAmount: number;
  creationDate: number;
  fundingDate: number | null;
  ttl_phase_one: number;
  ttl_phase_two: number;
  price: number;
  quantity: number;
  // timeseries: Timeseries;
  // orders: Order[];
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
