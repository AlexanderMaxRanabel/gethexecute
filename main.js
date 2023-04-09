import { run } from 'https://deno.land/std/command/mod.ts';

// Start Geth light client
const gethProcess = run({
  cmd: ['geth', '--syncmode=light', '--cache=1024', '--rpc', '--rpcaddr=localhost', '--rpcport=8545', '--rpcapi=eth,net,web3,personal,debug', '--ws', '--wsaddr=localhost', '--wsport=8546', '--wsapi=eth,net,web3,personal,debug', '--wsorigins=*', '--allow-insecure-unlock']
});

// Wait for Geth process to start
setTimeout(async () => {

  // Attach to Geth console
  const gethConsole = run({
    cmd: ['geth', 'attach', 'http://localhost:8545'],
    stdout: 'piped',
    stderr: 'piped'
  });

  // Forward console output to stdout
  for await (const output of gethConsole.stdout) {
    console.log(output);
  }
  for await (const error of gethConsole.stderr) {
    console.error(error);
  }

}, 5000);

// Forward console output to stdout
for await (const output of gethProcess.stdout) {
  console.log(output);
}
for await (const error of gethProcess.stderr) {
  console.error(error);
}
