using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SticksGame
{
    class RandomPlayer : Player
    {
        private static Random result = new Random();
        public RandomPlayer()
        {
            Name = "Random";
        }

        public override int GetMove(int sticksInPlay)
        {
            switch (sticksInPlay) {
                case 1:
                    return 1;
                default:
                    return result.Next(1, 4);
            }
        }
    }
}
