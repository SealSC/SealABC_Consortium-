/*
 * Copyright 2020 The SealABC Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package traceableStorage

import (
	"SealABC/crypto"
	"SealABC/crypto/hashes/sha3"
	"SealABC/crypto/signers/ed25519"
	commonCfg "SealABC/metadata/applicationCommonConfig"
	"SealABC/storage/db/dbDrivers/levelDB"
	"SealABC/storage/db/dbInterface"
)

type Config struct {
	commonCfg.Config

	CryptoTools crypto.Tools
}

func DefaultConfig() *Config {
	return &Config {
		Config: commonCfg.Config {
			KVDBName: dbInterface.LevelDB,
			KVDBConfig: levelDB.Config{
				DBFilePath: "./traceableStorage",
			},

			EnableSQLDB:  false,
			SQLStorage:   nil,
		},

		CryptoTools:  crypto.Tools{
			HashCalculator:  sha3.Sha256,
			SignerGenerator: ed25519.SignerGenerator,
		},
	}
}
