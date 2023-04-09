use std::process::Command;

fn main() {
    // Start Geth light client
    let geth_process = Command::new("geth")
        .arg("--syncmode=light")
        .arg("--cache=1024")
        .arg("--rpc")
        .arg("--rpcaddr=localhost")
        .arg("--rpcport=8545")
        .arg("--rpcapi=eth,net,web3,personal,debug")
        .arg("--ws")
        .arg("--wsaddr=localhost")
        .arg("--wsport=8546")
        .arg("--wsapi=eth,net,web3,personal,debug")
        .arg("--wsorigins=*")
        .arg("--allow-insecure-unlock")
        .spawn()
        .expect("Failed to start Geth process");

    // Wait for Geth process to start
    std::thread::sleep(std::time::Duration::from_secs(5));

    // Attach to Geth console
    let _geth_console = Command::new("geth")
        .arg("attach")
        .arg("http://localhost:8545")
        .spawn()
        .expect("Failed to attach to Geth console");

    // Wait for Geth process to exit
    let _result = geth_process.wait();
}
