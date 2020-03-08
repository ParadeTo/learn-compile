import repl from './repl'
import interpret from './interprete'

const argv = process.argv

// node you -v demo.you
// node you demo.you
// node you -v
// node you
if (argv.length >= 4) {
  if (argv[2] !== '-v') throw new Error('only support -v')
  interpret(argv[3], true)
} else if (argv.length === 3) {
  if (argv[2] === '-v') repl(true)
  else interpret(argv[2])
} else {
  repl()
}
