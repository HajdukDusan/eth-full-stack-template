import { ethers } from "ethers";

export let provider
export let chainId

export let signer
export let walletAddress
export let currentBalance

export async function connect() {
    if (window.ethereum) {
        provider = new ethers.providers.Web3Provider(window.ethereum);

        chainId = (await provider.getNetwork()).chainId;

        provider.send("eth_requestAccounts", [])
            .then(async () => {

                signer = provider.getSigner();
                walletAddress = await signer.getAddress()
                currentBalance = await signer.getBalance()
            })
            .catch(error => { throw new Error("Request Accounts Failed") });

    } else {
        throw new Error("Please Install Metamask.");
    }
}

export async function disconnect() {
    provider = null
    chainId = null
    walletAddress = null
    currentBalance = null
}

export async function contractView(address, abi, provider, ...callArgs) {
    const contract = new ethers.Contract(
        address,
        abi,
        provider
    );

    const result = await contract.stupidRegistry(
        ...callArgs
    );

    alert(result);
}

export async function contractTx(address, abi, ethValue, funcToCall, ...funcArgs) {

    const contractWithSigner = new ethers.Contract(
        address,
        abi,
        signer
    );

    const txOptions = { value: ethValue };

    const tx = await funcToCall(...funcArgs, txOptions)

    // let tx = await contractWithSigner.AddToRegistry(
    //     txMsgInputRef.current.value,
    //     txOptions
    // );
    alert(tx.hash);

    const receipt = await tx.wait();
    if (receipt.status !== 1) {
        alert("Tx failed!");
        return;
    }


}

function parseLogs(logs, abi) {
    let abiInterface = new ethers.utils.Interface(abi);
    return logs.map((log) => abiInterface.parseLog(log));

    //events[0].args[0]
    //events[0].name
    //events[0].signature
}

