/****************************************************
Copyright 2018 The ont-eventbus Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*****************************************************/

/***************************************************
Copyright 2016 https://github.com/AsynkronIT/protoactor-go

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*****************************************************/
package main

import (
	"runtime"
	"time"

	"github.com/ontio/ontology-eventbus/actor"
	"github.com/ontio/ontology-eventbus/eventhub"
	"github.com/ontio/ontology-eventbus/example/testRemoteCrypto/commons"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 1)
	runtime.GC()

	props := actor.FromProducer(func() actor.Actor { return &commons.BusynessActor{Datas: make(map[string][]byte)} })
	bActor := actor.Spawn(props)

	signprops := actor.FromProducer(func() actor.Actor { return &commons.SignActor{} })
	signActor := actor.Spawn(signprops)

	eventhub.GlobalEventHub.Subscribe(commons.SetTOPIC, signActor)
	eventhub.GlobalEventHub.Subscribe(commons.SigTOPIC, signActor)

	vfprops := actor.FromProducer(func() actor.Actor { return &commons.VerifyActor{} })
	vfActor := actor.Spawn(vfprops)

	eventhub.GlobalEventHub.Subscribe(commons.VerifyTOPIC, vfActor)

	bActor.Tell(&commons.RunMsg{})

	for {
		time.Sleep(1 * time.Second)
	}
}
