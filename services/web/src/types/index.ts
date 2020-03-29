export interface Superhero {
  id: number;
  fullName: string;
  alterEgo: string;
  imageUrl: string;
  intelligence: number;
  strength: number;
  speed: number;
  durability: number;
  power: number;
  combat: number;
}

export enum PlayerTurn {
  PlayerOne = 0,
  PlayerTwo
}

export enum Winner {
  None = 0,
  PlayerOne,
  PlayerTwo
}
