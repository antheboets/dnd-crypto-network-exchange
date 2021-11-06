using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace WebApplication2.Models
{
    public class Transaction
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public string Id { get; set; }
        public string TokenId { get; set; }
        public Token Token { get; set; }
        public double Amount { get; set; }
        public Wallet SenderWallet { get; set; }
        public string SenderWalletId { get; set; }
        public Wallet RecieverWallet { get; set; }
        public string RecieverWalletId { get; set; }
    }
}
