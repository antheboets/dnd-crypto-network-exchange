using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace WebApplication2.Models
{
    public class Wallet
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public string Id { get; set; }
        public double Value { get; set; }
        public string UserId { get; set; }
        public User User { get; set; }
        public Token Token { get; set; }
        public string TokenId { get; set; }
        public string WalletLink { get; set; }

    }
}
