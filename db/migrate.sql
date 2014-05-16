/* Migrating table names for clarity */

-- Users
rename table C_USERS to Users;
-- Articles
rename table C_ARTICLES to Articles;
-- Interviews
rename table C_INTERVIEWS to Interviews;
rename table F_CERTIFICATION to Interviews_Certification;
rename table L_CERTIFICATION to Interviews_Certification_Lookup;
rename table F_SALARY to Interviews_Salary;
rename table L_SALARY to Interviews_Salary_Lookup;
rename table F_SECTOR to Interviews_Sector;
rename table L_SECTOR to Interviews_Sector_Lookup;
rename table F_OVERNIGHT_TRAVEL to Interviews_Overnight_Travel;
rename table L_OVERNIGHT_TRAVEL to Interviews_Overnight_Travel_Lookup;
rename table F_TYPE to Interviews_Type;
rename table L_TYPE to Interviews_Type_Lookup;
rename table F_WEEKENDS_WORKED to Interviews_Weekends_Worked;
rename table L_WEEKENDS_WORKED to Interviews_Weekends_Worked_Lookup;
rename table F_EDUCATION to Interviews_Education;
rename table F_INDUSTRY to Interviews_Industry;
rename table L_INDUSTRY to Interviews_Industry_Lookup;
rename table F_DAILY_BREAKDOWN to Interviews_Daily_Breakdown;
rename table F_ENVIRONMENT to Interviews_Environment;
rename table L_ENVIRONMENT to Interviews_Environment_Lookup;
rename table F_EXPERIENCE to Interviews_Experience;
rename table L_EXPERIENCE to Interviews_Experience_Lookup;
rename table F_FULFILLMENT to Interviews_Fulfillment;
rename table F_HOURS_PER_WEEK to Interviews_Hours_Per_Week;
rename table L_HOURS_PER_WEEK to Interviews_Hours_Per_Week_Lookup;
rename table F_SKILL to Interviews_Skill;
rename table F_SOFT_SKILL to Interviews_Soft_Skills;
rename table L_SOFT_SKILL to Interviews_Soft_Skills_Lookup;
rename table F_SOLO_GROUP to Interviews_Groupwork;
rename table F_TOOL to Interviews_Tools;
rename table L_DEGREE to Interviews_Degree_Lookup;
rename table L_MAJOR to Interviews_Major_Lookup;
rename table L_UNIVERSITY to Interviews_University_Lookup;
