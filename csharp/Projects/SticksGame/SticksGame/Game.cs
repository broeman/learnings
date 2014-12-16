using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SticksGame
{
    class Game
    {
        private int numGames = 0;
        private Player[] players = new Player[2];
        private Stick sticks;
        private bool textInterface = false;

        public void Init()
        {
            Console.WriteLine(" ============ Velkommen til Pindespillet ============== \n");
            Console.WriteLine("1: Player");
            Console.WriteLine("2: Random");
            Console.WriteLine("3: Smart");
            Console.WriteLine("4: Human");
            Console.WriteLine("5: AI\n");

            int player1 = 0, player2 = 0;
            bool retry = true;

            while (retry)
            {
                try
                {
                    Console.Write("Vælg Player 1: ");
                    player1 = Convert.ToInt32(Console.ReadLine());
                    if (player1 < 0 || player1 > 5)
                        break;
                    else 
                    {
                        Console.Write("Vælg Player 2: ");
                        player2 = Convert.ToInt32(Console.ReadLine());
                        if (player2 > 0 || player2 < 6)
                            retry = false;
                    }                    
                }
                catch (FormatException)
                {
                    Console.WriteLine("Det skal være et tal");
                }
                catch (OverflowException)
                {
                    Console.WriteLine("It cannot fit in an Integer 32-bit");
                }

            }
            players[0] = ChoosePlayer(player1);
            players[1] = ChoosePlayer(player2);
            
        }

        private Player ChoosePlayer(int choice)
        {
            switch (choice)
            {
                case 1:
                    return new Player();
                case 2:
                    return new RandomPlayer();
                case 3:
                    return new SmartPlayer();
                case 4:
                    return new HumanPlayer();
                case 5:
                    return new IntelligentPlayer();
                default:
                    return new Player();
            }
        }

        public void Play() {
            sticks = new Stick(15);
            numGames++;

            for (int i = 0; i < players.Length; i++)
            {
                players[i].StartGame();
                if (players[i].PlayerType == "Human")
                {
                    textInterface = true;
                }
            }

            while (sticks.Count > 0)
            {
                for (int i = 0; i < players.Length; i++ )
                {
                    if (textInterface) {
                        Console.WriteLine("");
                        for (int j = 0; j < sticks.Count; j++)
                            Console.Write("|");
                        Console.Write(" ({0}) \n", sticks.Count);
                    }
                    int curSticks = sticks.Count;
                    int taken = players[i].GetMove(sticks.Count);
                    sticks.Count = curSticks - taken;
                    if (textInterface)
                        Console.WriteLine("{0} tager {1}", players[i].Name, taken);
                    if (sticks.Count <= 1)
                    {
                        for (int j = 0; j < players.Length; j++)
                        {
                            if (j == i)
                            {
                                players[j].EndGame(true);
                                if (textInterface)
                                {
                                    Console.WriteLine("{0} vandt!", players[j].Name);
                                    Console.WriteLine(String.Format("\n{0}. spil -- {1}: {2}  /  {3}: {4} --\n", numGames.ToString(), players[0].Name, players[0].Wins, players[1].Name, players[1].Wins));
                                }
                            }
                            else
                            {
                                players[j].EndGame(false);
                            }
                        }          
                        return;
                    }
                }
            }
        }
        public void Play(int gameCount) {
            bool running = true;
            while (running)
            {
                for (int i = 0; i < gameCount; i++)
                {
                    Play();
                }
                Console.WriteLine(GetStat());
                Console.WriteLine("Vil du spille {0} spil igen (j/n)? ", gameCount);
                string feedback = Console.ReadLine();
                if (feedback == "n" || feedback == "N")
                {
                    running = false;
                }

            }
        } 
        public string GetStat()
        {
            return String.Format("==== {0} spil, {1}: {2}  /  {3}: {4} ====", numGames.ToString(), players[0].Name, players[0].Wins, players[1].Name, players[1].Wins);
        }
    }
}
