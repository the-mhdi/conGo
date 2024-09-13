package opcode

var (
	STOP       = "00" //0
	ADD        = "01" //3	a, b => a + b		(u)int256 addition modulo 2**256
	MUL        = "02" //5	a, b => a * b		(u)int256 multiplication modulo 2**256
	SUB        = "03" //3	a, b => a - b		(u)int256 addition modulo 2**256
	DIV        = "04" //5	a, b => a // b		uint256 division
	SDIV       = "05" //5	a, b => a // b		int256 division
	MOD        = "06" //5	a, b => a % b		uint256 modulus
	SMOD       = "07" //5	a, b => a % b		int256 modulus
	ADDMOD     = "08" //8	a, b, N => (a + b) % N		(u)int256 addition modulo N
	MULMOD     = "09" //8	a, b, N => (a * b) % N		(u)int256 multiplication modulo N
	EX         = "0A" //A1	a, b => a ** b		uint256 exponentiation modulo 2**256
	SIGNEXTEND = "0B" //5		b, x => SIGNEXTEND(x, b)		sign extend x from (b+1) bytes to 32 bytes
	//0C-0F	invalid
	LT     = "10" //3	a, b => a < b		uint256 less-than
	GT     = "11" //3	a, b => a > b		uint256 greater-than
	SLT    = "12" //3	a, b => a < b		int256 less-than
	SGT    = "13" //3	a, b => a > b		int256 greater-than
	EQ     = "14" //3	a, b => a == b		(u)int256 equality
	ISZERO = "15" //3	a => a == 0		(u)int256 iszero
	AND    = "16" //3	a, b => a && b		bitwise AND
	OR     = "17" //3	a, b => a || b		bitwise OR
	XOR    = "18" //3	a, b => a ^ b		bitwise XOR
	NOT    = "19" //3	a => ~a		bitwise NOT
	BYTE   = "1A" //3	i, x => (x >> (248 - i * 8)) && 0xFF		ith byte of (u)int256 x, from the left
	SHL    = "1B" //3	shift, val => val << shift		shift left
	SHR    = "1C" //3	shift, val => val >> shift		logical shift right
	SAR    = "1D" //3	shift, val => val >> shift		arithmetic shift right
	//1E-1F	invalid
	SHA3 = "20" //A2	ost, len => keccak256(mem[ost:ost+len])		keccak256
	//21-2F	invalid
	ADDRESS        = "30" //2	. => address(this)		address of executing contract
	BALANCE        = "31" //A5	addr => addr.balance		balance, in wei
	ORIGIN         = "32" //2	. => tx.origin		address that originated the tx
	CALLER         = "33" //2	. => msg.sender		address of msg sender
	CALLVALUE      = "34" //2	. => msg.value		msg value, in wei
	CALLDATALOAD   = "35" //3	idx => msg.data[idx:idx+32]		read word from msg data at index idx
	CALLDATASIZE   = "36" //2	. => len(msg.data)		length of msg data, in bytes
	CALLDATACOPY   = "37" //A3	dstOst, ost, len => .	mem[dstOst:dstOst+len] := msg.data[ost:ost+len	copy msg data
	CODESIZE       = "38" //2	. => len(this.code)		length of executing contract's code, in bytes
	CODECOPY       = "39" //A3	dstOst, ost, len => .		mem[dstOst:dstOst+len] := this.code[ost:ost+len]
	GASPRICE       = "3A" //2	. => tx.gasprice		gas price of tx, in wei per unit gas **
	EXTCODESIZE    = "3B" //A5	addr => len(addr.code)		size of code at addr, in bytes
	EXTCODECOPY    = "3C" //A4	addr, dstOst, ost, len => .	mem[dstOst:dstOst+len] := addr.code[ost:ost+len]	copy code from addr
	RETURNDATASIZE = "3D" //2	. => size		size of returned data from last external call, in bytes
	RETURNDATACOPY = "3E" //A3	dstOst, ost, len => .	mem[dstOst:dstOst+len] := returndata[ost:ost+len]	copy returned data from last external call
	EXTCODEHASH    = "3F" //A5	addr => hash		hash = addr.exists ? keccak256(addr.code) : 0
	BLOCKHASH      = "40" //20	blockNum => blockHash(blockNum)
	COINBASE       = "41" //2	. => block.coinbase		address of miner of current block
	TIMESTAMP      = "42" //2	. => block.timestamp		timestamp of current block
	NUMBER         = "43" //2	. => block.number		number of current block
	DIFFICULTY     = "44" //2	. => block.difficulty		difficulty of current block
	GASLIMIT       = "45" //2	. => block.gaslimit		gas limit of current block
	CHAINID        = "46" //2	. => chain_id		push current chain id onto stack
	SELFBALANCE    = "47" //5	. => address(this).balance		balance of executing contract, in wei
	BASEFEE        = "48" //2	. => block.basefee		base fee of current block
	//49-4F	invalid
	POP      = "50" //2	_anon => .		remove item from top of stack and discard it
	MLOAD    = "51" //3*	ost => mem[ost:ost+32]		read word from memory at offset ost
	MSTORE   = "52" //3*	ost, val => .	mem[ost:ost+32] := val	write a word to memory
	MSTORE8  = "53" //3*	ost, val => .	mem[ost] := val && 0xFF	write a single byte to memory
	SLOAD    = "54" //A6	key => storage[key]		read word from storage
	SSTORE   = "55" //A7	key, val => .	storage[key] := val	write word to storage
	JUMP     = "56" //8	dst => .		$pc := dst
	JUMPI    = "57" //10	dst, condition => .		$pc := condition ? dst : $pc + 1
	PC       = "58" //2	. => $pc		program counter
	MSIZE    = "59" //2	. => len(mem)		size of memory in current execution context, in bytes
	GAS      = "5A" //2	. => gasRemaining
	JUMPDEST = "5B" //1			mark valid jump destination
	//5C-5F	invalid
	PUSH1  = "60" //3	. => uint8			push 1-byte value onto stack
	PUSH2  = "61" //3	. => uint16			push 2-byte value onto stack
	PUSH3  = "62" //3	. => uint24			push 3-byte value onto stack
	PUSH4  = "63" //3	. => uint32			push 4-byte value onto stack
	PUSH5  = "64" //3	. => uint40			push 5-byte value onto stack
	PUSH6  = "65" //3	. => uint48			push 6-byte value onto stack
	PUSH7  = "66" //3	. => uint56			push 7-byte value onto stack
	PUSH8  = "67" //3	. => uint64			push 8-byte value onto stack
	PUSH9  = "68" //3	. => uint72			push 9-byte value onto stack
	PUSH10 = "69" //3	. => uint80			push 10-byte value onto stack
	PUSH11 = "6A" //3	. => uint88			push 11-byte value onto stack
	PUSH12 = "6B" //3	. => uint96			push 12-byte value onto stack
	PUSH13 = "6C" //3	. => uint104		push 13-byte value onto stack
	PUSH14 = "6D" //3	. => uint112		push 14-byte value onto stack
	PUSH15 = "6E" //3	. => uint120		push 15-byte value onto stack
	PUSH16 = "6F" //3	. => uint128		push 16-byte value onto stack
	PUSH17 = "70" //3	. => uint136		push 17-byte value onto stack
	PUSH18 = "71" //3	. => uint144		push 18-byte value onto stack
	PUSH19 = "72" //3	. => uint152		push 19-byte value onto stack
	PUSH20 = "73" //3	. => uint160		push 20-byte value onto stack
	PUSH21 = "74" //3	. => uint168		push 21-byte value onto stack
	PUSH22 = "75" //3	. => uint176		push 22-byte value onto stack
	PUSH23 = "76" //3	. => uint184		push 23-byte value onto stack
	PUSH24 = "77" //3	. => uint192		push 24-byte value onto stack
	PUSH25 = "78" //3	. => uint200		push 25-byte value onto stack
	PUSH26 = "79" //3	. => uint208		push 26-byte value onto stack
	PUSH27 = "7A" //3	. => uint216		push 27-byte value onto stack
	PUSH28 = "7B" //3	. => uint224		push 28-byte value onto stack
	PUSH29 = "7C" //3	. => uint232		push 29-byte value onto stack
	PUSH30 = "7D" //3	. => uint240		push 30-byte value onto stack
	PUSH31 = "7E" //3	. => uint248		push 31-byte value onto stack
	PUSH32 = "7F" //3	. => uint256		push 32-byte value onto stack
	DUP1   = "80" //3	a => a, a			clone 1st value on stack
	DUP2   = "81" //3	_, a => a, _, a		clone 2nd value on stack
	DUP3   = "82" //3	_, _, a => a, _, _, a		clone 3rd value on stack
	DUP4   = "83" //3	_, _, _, a => a, _, _, _, a		clone 4th value on stack
	DUP5   = "84" //3	..., a => a, ..., a		clone 5th value on stack
	DUP6   = "85" //3	..., a => a, ..., a		clone 6th value on stack
	DUP7   = "86" //3	..., a => a, ..., a		clone 7th value on stack
	DUP8   = "87" //3	..., a => a, ..., a		clone 8th value on stack
	DUP9   = "88" //3	..., a => a, ..., a		clone 9th value on stack
	DUP10  = "89" //3	..., a => a, ..., a		clone 10th value on stack
	DUP11  = "8A" //3	..., a => a, ..., a		clone 11th value on stack
	DUP12  = "8B" //3	..., a => a, ..., a		clone 12th value on stack
	DUP13  = "8C" //3	..., a => a, ..., a		clone 13th value on stack
	DUP14  = "8D" //3	..., a => a, ..., a		clone 14th value on stack
	DUP15  = "8E" //3	..., a => a, ..., a		clone 15th value on stack
	DUP16  = "8F" //3	..., a => a, ..., a		clone 16th value on stack
	SWAP1  = "90" //3	a, b => b, a
	SWAP2  = "91" //3	a, _, b => b, _, a
	SWAP3  = "92" //3	a, _, _, b => b, _, _, a
	SWAP4  = "93" //3	a, _, _, _, b => b, _, _, _, a
	SWAP5  = "94" //3	a, ..., b => b, ..., a
	SWAP6  = "95" //3	a, ..., b => b, ..., a
	SWAP7  = "96" //3	a, ..., b => b, ..., a
	SWAP8  = "97" //3	a, ..., b => b, ..., a
	SWAP9  = "98" //3	a, ..., b => b, ..., a
	SWAP10 = "99" //3	a, ..., b => b, ..., a
	SWAP11 = "9A" //3	a, ..., b => b, ..., a
	SWAP12 = "9B" //3	a, ..., b => b, ..., a
	SWAP13 = "9C" //3	a, ..., b => b, ..., a
	SWAP14 = "9D" //3	a, ..., b => b, ..., a
	SWAP15 = "9E" //3	a, ..., b => b, ..., a
	SWAP16 = "9F" //3	a, ..., b => b, ..., a
	LOG0   = "A0" //A8	ost, len => .		LOG0(memory[ost:ost+len])
	LOG1   = "A1" //A8	ost, len, topic0 => .		LOG1(memory[ost:ost+len], topic0)
	LOG2   = "A2" //A8	ost, len, topic0, topic1 => .		LOG1(memory[ost:ost+len], topic0, topic1)
	LOG3   = "A3" //A8	ost, len, topic0, topic1, topic2 => .		LOG1(memory[ost:ost+len], topic0, topic1, topic2)
	LOG4   = "A4" //A8	ost, len, topic0, topic1, topic2, topic3 => .		LOG1(memory[ost:ost+len], topic0, topic1, topic2, topic3)
	//A5-EF	invalid
	CREATE       = "F0" //A9	val, ost, len => addr		addr = keccak256(rlp_encode([address(this), this.nonce]))
	CALL         = "F1" //AA	gas, addr, val, argOst, argLen, retOst, retLen => success	mem[retOst:retOst+retLen] := returndata
	CALLCODE     = "F2" //AA	gas, addr, val, argOst, argLen, retOst, retLen => success	mem[retOst:retOst+retLen] = returndata	same as DELEGATECALL, but does not propagate original msg.sender and msg.value
	RETURN       = "F3" //0*	ost, len => .		return mem[ost:ost+len]
	DELEGATECALL = "F4" //AA	gas, addr, argOst, argLen, retOst, retLen => success		mem[retOst:retOst+retLen] := returndata
	CREATE2      = "F5" //A9	val, ost, len, salt => addr		addr = keccak256(0xff ++ address(this) ++ salt ++ keccak256(mem[ost:ost+len]))[12:]
	//F6-F9	invalid
	STATICCALL = "FA" //AA	gas, addr, argOst, argLen, retOst, retLen => success	mem[retOst:retOst+retLen] := returndata
	//FB-FC	invalid
	REVERT       = "FD" //0*	ost, len => .		revert(mem[ost:ost+len])
	INVALID      = "FE" //AF			designated invalid opcode - EIP-141
	SELFDESTRUCT = "FF" //AB	addr => .
)
