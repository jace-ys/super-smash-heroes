export interface Superhero {
  id: number;
  full_name: string;
  alter_ego: string;
  image_url: string;
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
