// Copyright 2022-2024 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bpfprograms

const (
	StackDepth       = 127 // Always needs to be sync with MAX_STACK_DEPTH in BPF program.
	tripleStackDepth = StackDepth * 3
)

var (

	// native programs.
	NativeProgramFD           = uint64(0)
	RubyEntrypointProgramFD   = uint64(1)
	PythonEntrypointProgramFD = uint64(2)
	// rbperf programs.
	RubyUnwinderProgramFD = uint64(0)
	// python programs.
	PythonUnwinderProgramFD = uint64(0)

	ProgramName               = "entrypoint"
	NativeUnwinderProgramName = "native_unwind"
)

type CombinedStack [tripleStackDepth]uint64
