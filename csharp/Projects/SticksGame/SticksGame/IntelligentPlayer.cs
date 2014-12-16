using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SticksGame
{
    class IntelligentPlayer : RandomPlayer
    {        
        private int curMove, numPlay;
        private Dictionary<Int32, Int32> moves = null;
        private List<List<int>> pluslist = new List<List<int>>();
        private List<List<int>> minuslist = new List<List<int>>();
        public IntelligentPlayer()
        {
            Name = "EnSteen";
            for (int i = 0; i < 15; i++)
                for (int j = 0; j < 3; j++)
                    moves[i, j] = 0;

        }

        public override void StartGame()
        {
            numPlay = 0;
            curMove = 15;
            moves = new Dictionary<int,int>();
        }

        public override int GetMove(int sticksInPlay)
        {
            int otherMove = curMove - sticksInPlay;
            curMove -= sticksInPlay;
            if (curMove != 0) 
            {
                numPlay++;
                //moves.Add();
            }

            if (pluslist != null)
            {
                foreach (List<int> list in pluslist)
                {
                    if (list[numPlay] == otherMove)
                    {
                        int bestTry = list[numPlay + 1];
                        //moves.Add(bestTry);
                        curMove = sticksInPlay - bestTry;
                        return bestTry;
                    }
                    else
                    {
                        break;
                    }

                }
            }
            //

            var baseTry = base.GetMove(sticksInPlay);
            //moves.Add(baseTry);
            curMove = sticksInPlay - baseTry;
            return baseTry;
        }

        public override void EndGame(bool youHaveWon)
        {
            if (youHaveWon) {                
                //pluslist.Add(moves);
                base.EndGame(youHaveWon);
            }
            else
            {
                //minuslist.Add(moves);
            }
        }
    }
}
