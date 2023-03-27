CC := g++

# Флаги компиляции
CFLAGS := -std=c++17

# Папки
SRCDIR := src/
INCDIR := inc/
OUTDIR := out/

# Цель сборки
TARGET := app.exe

SRCS := $(wildcard $(SRCDIR)*.cpp)
OBJS := $(SRCS:$(SRCDIR)%.cpp=$(OUTDIR)%.o)
DEPS := $(SRCS:$(SRCDIR)%.cpp=$(OUTDIR)%.d)

.PHONY : debug
debug : $(TARGET)

.PHONY : release
release : $(TARGET)

$(TARGET) : $(OBJS)
	$(CC) $(CFLAGS) $^ -o $@ -lm -lpthread

$(OUTDIR)%.o : $(SRCDIR)%.cpp | $(OUTDIR)
	$(CC) $(CFLAGS) -I$(INCDIR) -MD -o $@ -c $<

-include $(DEPS)

$(OUTDIR) :
	mkdir $@

.PHONY : clean
clean :
	@rm -f $(TARGET) $(OUTDIR)*
