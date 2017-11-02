#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>
#include <string.h>
#include <assert.h>
void printRgb(char *string, int r, int g, int b){
		printf("\x1b[38;2;%i;%i;%im%s\x1b[0m",r,g,b,string);
}
bool fileExists(const char * file_name){
		FILE *file;
		if(file = fopen(file_name, "r")) {
				fclose(file);
				return 1;
		}
		return 0;


}
char *getDate(){

		time_t now;
		time(&now);

		struct tm* now_tm;
		now_tm = localtime(&now);
		char *date=malloc(sizeof(date));
		strftime (date, 12, "%m-%d-%y", now_tm);
		return date;
}
char *appendFile(char *file_name,char *text){
		FILE *file=fopen(file_name, "a");
		fputs(text,file);
		fclose(file);
}
char subjects[][100]={"biology","englisdfsdfsdfsh","history","math","spanish"};

void nothing(char *the_file){
		char *s = (char*) malloc( 100 );
		char *method="a";
		if(fileExists(the_file)) {
				char choice;
				printf("Already plannerized, are you sure you want to reset (y/n):");
				scanf("%c",&choice);
				if(choice=='y') {
						method="w";
				}
				else{
						return;
				}
		}
		FILE *file=fopen(the_file, method);
		for(int i = 0; i < sizeof(subjects)/sizeof(subjects[0]); i++) {
				printf("Homework for %s: ", subjects[i]);
				scanf("%s", s);
				fputs(subjects[i],file);
				fputs(": ",file);
				fputs(s,file);

				fputs("\n",file);

		}
		fclose(file);
		return;
}
void list(char *the_file){
		FILE *file=fopen(the_file,"r");
		char str[200];

		while(fgets(str,200,file)!=NULL) {
				if((strstr(str, "done") != NULL) || (strstr(str, "none") != NULL)) {
						printRgb(str,50,255,100 );
				}
				else{
						printRgb(str,255, 50, 50);
				}
		}
		fclose(file);
}
void edit(char * subject,char *the_file){
		char *s = (char*) malloc( 100 );
		printf("Homework for %s: ",subject);
		scanf("%s", s);
		printf("%s\n",s );
		FILE *file=fopen(the_file,"r");
		FILE *tmp=fopen("tmp","w");
		char str[200];

		while(fgets(str,200,file)!=NULL) {
				if((strstr(str, subject) != NULL)) {
						fputs(subject,tmp);
						fputs(": ",tmp);
						fputs(s,tmp);
						fputs("\n",tmp);
				}
				else{
						fputs(str,tmp);
				}

		}
		fclose(file);
		fclose(tmp);

}


int main(){

		// nothing(file);
		// list(file);
		edit("history","/home/god/Documents/school/planner/11-02-17");




		// printf("%s\n",new_folder );

}
