class ATMCard:
    # Initialize PIN and Balance
    def __init__(self, defaultPin, defaultBalance):
        self.defaultPin = defaultPin
        self.defaultBalance = defaultBalance
    # Validate PIN of ATM
    def cekPinAwal(self):
        return self.defaultPin
    # Show current balance
    def cekSaldoAwal(self):
        return self.defaultBalance
