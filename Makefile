#初始化项目目录变量
HOMEDIR := $(shell pwd)
OUTDIR  := $(HOMEDIR)/output

# 可以修改为其他的名字
APPNAME = $(shell basename `pwd`)
OUTPUT_FILE=${APPNAME}.tar.gz

#初始化命令变量
GOROOT  := /usr/local/go
GO      := $(GOROOT)/bin/go
GOPATH  := $(shell $(GO) env GOPATH)
GOMOD   := $(GO) mod
GOBUILD := $(GO) build
GOTEST  := $(GO) test
GOPKGS  := $$($(GO) list ./...| grep -vE "vendor")

SCRIPT_LIST := $(shell cd $(HOMEDIR) && ls scripts/erg3020/*/*.go && cd $(HOMEDIR))
SCRIPT_TARGET := $(SCRIPT_LIST:%.go=%)

#执行编译，可使用命令 make 或 make all 执行, 顺序执行prepare -> compile -> test -> package 几个阶段
all: compile package
all-offline: compile package-offline

#complile阶段，执行编译命令,可单独执行命令: make compile
compile:build

build: 
	$(GOMOD) download #下载Go依赖
	$(GOBUILD) -o $(HOMEDIR)/bin/$(APPNAME)
	$(shell cd $(HOMEDIR) && rm -f $(SCRIPT_TARGET) && cd $(HOMEDIR))
	@for target in $(SCRIPT_TARGET); do \
		$(GOBUILD) -o $(HOMEDIR)/$$target $$target.go ; \
	done

#test阶段，进行单元测试， 可单独执行命令: make test
test: test-case
test-case: set-env
	$(GOTEST) -race -v -cover $(GOPKGS) -gcflags="-N -l"

#package阶段，对编译产出进行打包，输出到output目录, 可单独执行命令: make package
package: package-bin
package-bin:
	$(shell rm -rf $(OUTDIR))
	$(shell mkdir -p $(OUTDIR))
	$(shell cp -a bin $(OUTDIR)/bin)
	$(shell cp -a conf $(OUTDIR)/conf)
	$(shell cp -a scripts $(OUTDIR)/scripts)
	$(shell if [ -d "data_online"  ]; then cp -r data_online $(OUTDIR)/data; fi)
	$(shell cd $(OUTDIR)/; tar -zcf ${OUTPUT_FILE} ./*; rm -rf bin conf supervise data scripts)

package-offline: package-bin-offline
package-bin-offline:
	$(shell rm -rf $(OUTDIR))
	$(shell mkdir -p $(OUTDIR))
	$(shell cp -a bin $(OUTDIR)/bin)
	$(shell cp -a conf_offline $(OUTDIR)/conf)
	$(shell cp -a scripts $(OUTDIR)/scripts)
	$(shell if [ -d "data"  ]; then cp -r data $(OUTDIR)/data; fi)
	$(shell cd $(OUTDIR)/; tar -zcf ${OUTPUT_FILE} ./*; rm -rf bin conf supervise data scripts)

#clean阶段，清除过程中的输出, 可单独执行命令: make clean
clean:
	rm -rf $(OUTDIR)

# avoid filename conflict and speed up build
.PHONY: all prepare compile test package  clean build
