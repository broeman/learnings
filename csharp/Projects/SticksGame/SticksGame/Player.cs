using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SticksGame
{
    class Player
    {
        public string PlayerType { get; protected set; }
        public string Name { get; set; }
        private int _wins = 0;
        public int Wins
        {
            get { return _wins; }
            private set { _wins = value; }
        }

        public Player()
        {
            Name = "Simple";
            Wins = 0;
        }

        public string GetName()
        {
            return Name;
        }

        public virtual void StartGame() { }
        public virtual int GetMove(int sticksInPlay)
        {
            return 1;
        }

        public virtual void EndGame(bool youHaveWon)
        {
            if (youHaveWon) {
                _wins++;
            }
        }
    }
}
