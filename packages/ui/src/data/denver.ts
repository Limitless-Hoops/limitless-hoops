export type TeamFont =
  | 'font-team-staatliches'
  | 'font-team-skranji'
  | 'font-team-flamenco'
  | 'font-team-nanumBrushScript'
  | 'font-team-comicNeue';

export interface TeamColors {
  primary: string;
  secondary: string | null;
}

export interface Team {
  id: string;
  name: string;
  location: string;
  city: string;
  moto: string;
  font: TeamFont;
  colors: TeamColors;
  emoji: string;
}

export const denverTeams: Team[] = [
  {
    id: 'bucks',
    name: 'Bucks',
    location: 'Denver',
    city: 'Denver',
    moto: 'RULES OF LIFE: HUNT OR BE HUNTED',
    font: 'font-team-staatliches',
    colors: {
      primary: '#eab47f',
      secondary: '#6b4423',
    },
    emoji: '🦌',
  },
  {
    id: 'bulls',
    name: 'Bulls',
    location: 'Broomfield',
    city: 'Broomfield',
    moto: 'A STRONG BULL IS SEEN BY ITS SCARS',
    font: 'font-team-staatliches',
    colors: {
      primary: '#cb6649',
      secondary: '#71706e',
    },
    emoji: '🐂',
  },
  {
    id: 'dragons',
    name: 'Dragons',
    location: 'Centennial',
    city: 'Centennial',
    font: 'font-team-flamenco',
    moto: "LIGHT A FIRE THEY CAN'T PUT OUT!",
    colors: {
      primary: '#71a6d2',
      secondary:'#50404d',
    },
    emoji: ' 🐲',
  },
  {
    id: 'eagles',
    name: 'Eagles',
    location: 'Denver',
    city: 'Denver',
    font: 'font-team-staatliches',
    moto: 'An EAGLE USES THE STORM TO REACH NEW HEIGHTS',
    colors: {
      primary: '#ffe4cd',
      secondary: '#351e1c',
    },
    emoji: ' 🦅',
  },
  {
    id: 'gators',
    name: 'Gators',
    location: 'Aurora',
    city: 'Aurora',
    font: 'font-team-staatliches',
    moto: "DON'T TAUNT THE GATOR... UNTIL AFTER YOU'VE CROSSED THE CREEK",
    colors: {
      primary: '#66b447',
      secondary: '#d7837f',
    },
    emoji: '🐊',
  },
  {
    id: 'hornets',
    name: 'Hornets',
    location: 'Aurora',
    city: 'Aurora',
    font: 'font-team-staatliches',
    moto: "STIR UP A HORNET'S NEST, NO TELLIN' WHO'S GON' GET STUNG!",
    colors: {
      primary: '#f8c176',
      secondary: '#783d3e',
    },
    emoji: ' 🐝',
  },
  {
    id: 'jokers',
    name: 'Jokers',
    location: 'Denver',
    city: 'Denver',
    font: 'font-team-comicNeue',
    moto: "HAHA!",
    colors: {
      primary: '#eb4c42',
      secondary: '#45b1e8',
    },
    emoji: '🃏',
  },
  {
    id: 'lions',
    name: 'Lions',
    location: 'Denver',
    city: 'Denver',
    font: 'font-team-staatliches',
    moto: "EVERYONE WANTS TO EAT, BUT FEW ARE WILLING TO HUNT",
    colors: {
      primary: '#ffc87c', 
      secondary: '#905d5d',
    },
    emoji: '🦁',
  },
  {
    id: 'pirates',
    name: 'Pirates',
    location: 'Centennial',
    city: 'Centennial',
    font: 'font-team-nanumBrushScript',
    moto: "WORK LIKE A CAPTION, PLAY LIKE A PIRATE",
    colors: {
      primary: '#d73b3e',
      secondary: '#d6d6d6',
    },
    emoji: '🏴‍☠️',
  },
  {
    id: 'raptors',
    name: 'Raptors',
    location: 'Broomfield',
    city: 'Broomfield',
    font: 'font-team-staatliches',
    moto: "FEROCITY IS AN ART ITSELF",
    colors: {
      primary: '#d7837f',
      secondary: '#4c516d',
    },
    emoji: '🦖'
  },
  {
    id: 'sharks',
    name: 'Sharks',
    location: 'Aurora',
    city: 'Aurora',
    font: 'font-team-staatliches',
    moto: "IN A WORLD FULL OF FISH BE A SHARK!",
    colors: {
      primary: '#006994',
      secondary: '#87cefa',
    },
    emoji: '🦈'
  },
  {
    id: 'stars',
    name: 'Stars',
    location: 'Broomfield',
    city: 'Broomfield',
    font: 'font-team-staatliches',
    moto: "DON'T JUST LOOK AT THE STARS - BE ONE",
    colors: {
      primary: '#ffc34d',
      secondary: '#000000',
    },
    emoji: '⭐'
  },
  {
    id: 'tigers',
    name: 'Tigers',
    location: 'Centennial',
    city: 'Centennial',
    font: 'font-team-staatliches',
    moto: "THERE'S NO OFF SWITCH ON A TIGER",
    colors: {
      primary: '#e88e5a',
      secondary: '#000000',
    },
    emoji: '🐯'
  },
  {
    id: 'vikings',
    name: 'Vikings',
    location: 'Aurora',
    city: 'Aurora',
    font: 'font-team-skranji',
    moto: "BETTER TO FIGHT AND FALL THAN LIVE WITHOUT HOPE",
    colors: {
      primary: '#9b533f',
      secondary: '#bfc1c2',
    },
    emoji: '🛡️',
  },
  {
    id: 'warriors',
    name: 'Warriors',
    location: 'Broomfield',
    city: 'Broomfield',
    font: 'font-team-staatliches',
    moto: "THE TWO MOST POWERFUL WARRIORS ARE PATIENCE AND TIME",
    colors: {
      primary: '#c53151',
      secondary: '#a9b2c3',
    },
    emoji: '⚔',
  },
  {
    id: 'wolves',
    name: 'Wolves',
    location: 'Centennial',
    city: 'Centennial',
    font: 'font-team-staatliches',
    moto: "IT NEVER TROUBLES THE WOLF HOW MANY SHEEP MAY BE...",
    colors: {
      primary: '#716675',
      secondary: '#a8a9ad',
    },
    emoji: '🐺'
  },
];