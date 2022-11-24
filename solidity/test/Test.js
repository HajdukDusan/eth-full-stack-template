const { ethers } = require("hardhat");
const { expect } = require("chai");

describe("PaymentContract", function () {

  let PaymentContract;

  const receivers = [
    "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
    "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
    "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
  ]
  const values = [200, 100, 250]

  before(async () => {
    const PaymentContractFactory = await ethers.getContractFactory("PaymentContract");
    PaymentContract = await PaymentContractFactory.deploy(
      "0x61ec9Cbc365b23eC035986A30FDed12e94756b3B"
    );
  });

  it("Create new payment subscription and pay", async function () {

    const currentDate = new Date();
    const startTime = currentDate.getTime();

    const subTx = await PaymentContract.CreatePaymentSubscriptionEth(
      startTime,
      10000,
      receivers,
      values,
      { value: ethers.utils.parseEther("1") }
    );
    await subTx.wait();

    const receiversBalanceBefore = []

    for (let i = 0; i < receivers.length; i++) {
      receiversBalanceBefore.push(await PaymentContract.provider.getBalance(receivers[i]))
    }

    const payTx = await PaymentContract.PayPaymentEth(0);
    await payTx.wait();

    for (let i = 0; i < receivers.length; i++) {
      expect(await PaymentContract.provider.getBalance(receivers[i]))
        .to.equal(receiversBalanceBefore[i].add(values[i]));
    }
  });

  it("Create new payment subscription and pay to the end", async function () {

    const currentDate = new Date();
    const startTime = currentDate.getTime();

    const subTx = await PaymentContract.CreatePaymentSubscriptionEth(
      startTime,
      10000,
      receivers,
      values,
      { value: 600 }
    );
    await subTx.wait();

    const receiversBalanceBefore = []

    for (let i = 0; i < receivers.length; i++) {
      receiversBalanceBefore.push(await PaymentContract.provider.getBalance(receivers[i]))
    }

    const payTx = await PaymentContract.PayPaymentEth(1);
    await payTx.wait();

    await expect(await PaymentContract.PayPaymentEth(1))
      .to.emit(PaymentContract, "PaymentSubscriptionEnded");
  });

  // it("Set new price for ETH symbol", async function () {

  //   const newPrice = ethers.utils.parseEther("1.1");

  //   const setPriceTx = await PriceSetter.set("ETH", newPrice);
  //   await setPriceTx.wait();

  //   const currentPrice = await PriceSetter.priceOf("ETH");

  //   expect(newPrice.toString()).to.equal(currentPrice.toString());
  // });

  // it("Try to set lower than 2% price difference", async function () {

  //   const setupPriceTx = await PriceSetter.set("ETH", ethers.utils.parseEther("1"));
  //   await setupPriceTx.wait();

  //   const newPrice = ethers.utils.parseEther("1.02");

  //   await expect(PriceSetter.set("ETH", newPrice)).to.be.revertedWith("PriceSetter__InsufficientPriceDifference()");
  // });

  // it("Try to set with empty symbol", async function () {
  //   await expect(PriceSetter.set("", ethers.utils.parseEther("1"))).to.be.revertedWith("PriceSetter__SymbolCantBeEmpty()");
  // });
});