using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SticksGame
{
    class Program
    {
        static void Main(string[] args)
        {            
            Game game;
            bool running = true;
            while (running)
            {
                game = new Game();
                game.Init();
                Console.Write("Antal spil? ");
                string numGames = Console.ReadLine();
                int numVal = 1;
                try
                {
                    numVal = Convert.ToInt32(numGames);
                    if (numVal < 1)
                    {
                        Console.WriteLine("Der skal mindst være ét spil.");
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
                game.Play(numVal);
                Console.WriteLine(game.GetStat());
                Console.WriteLine("Vil du spille igen (j/n)? ");
                string feedback = Console.ReadLine();
                if (feedback == "n" || feedback == "N")
                {
                    running = false;
                }
            }            

        }
    }
}
